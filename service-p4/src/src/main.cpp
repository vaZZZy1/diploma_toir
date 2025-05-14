#include <iostream>
#include <fstream>
#include <string>
#include <boost/asio.hpp>
#include <boost/program_options.hpp>
#include <nlohmann/json.hpp>
#include <memory>
#include <signal.h>

// Заглушки для будущих файлов
#include "server/Server.h"
#include "config/ConfigManager.h"
#include "logger/Logger.h"

using json = nlohmann::json;
namespace po = boost::program_options;

// Глобальные объекты для корректного завершения
std::shared_ptr<inventory::server::Server> server;
std::unique_ptr<inventory::logger::Logger> logger;

// Обработчик сигналов
void signalHandler(int signum) {
    if (logger) {
        logger->info("Получен сигнал: " + std::to_string(signum) + ", завершение работы...");
    }
    
    if (server) {
        server->stop();
    }
    
    exit(signum);
}

int main(int argc, char* argv[]) {
    // Настройка обработки сигналов
    signal(SIGINT, signalHandler);
    signal(SIGTERM, signalHandler);
    
    try {
        // Парсинг аргументов командной строки
        po::options_description desc("Доступные опции");
        desc.add_options()
            ("help", "Вывести справку")
            ("config", po::value<std::string>()->default_value("config/config.json"), "Путь к конфигурационному файлу");
        
        po::variables_map vm;
        po::store(po::parse_command_line(argc, argv, desc), vm);
        po::notify(vm);
        
        if (vm.count("help")) {
            std::cout << desc << std::endl;
            return 0;
        }
        
        // Загрузка конфигурации
        std::string configPath = vm["config"].as<std::string>();
        auto configManager = std::make_shared<inventory::config::ConfigManager>(configPath);
        
        // Инициализация логгера
        logger = std::make_unique<inventory::logger::Logger>(
            configManager->getLogLevel(),
            configManager->getLogFile()
        );
        
        logger->info("Сервис управления складом и потребностями запускается...");
        logger->info("Загружена конфигурация из: " + configPath);
        
        // Инициализация сервера
        server = std::make_shared<inventory::server::Server>(
            configManager,
            logger
        );
        
        // Запуск сервера
        server->start();
        
        // Основной цикл обработки сообщений
        logger->info("Сервис управления складом и потребностями запущен на порту " + 
                    std::to_string(configManager->getServerPort()));
        
        server->waitForTermination();
        
        logger->info("Сервис управления складом и потребностями завершил работу");
        return 0;
    } 
    catch (const std::exception& e) {
        if (logger) {
            logger->error("Критическая ошибка: " + std::string(e.what()));
        } else {
            std::cerr << "Критическая ошибка: " << e.what() << std::endl;
        }
        return 1;
    }
} 