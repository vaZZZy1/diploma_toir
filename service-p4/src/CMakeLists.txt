cmake_minimum_required(VERSION 3.10)
project(inventory_service VERSION 1.0.0)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra")

# Включаем поддержку тестирования
enable_testing()

# Находим необходимые библиотеки
find_package(Boost 1.71 REQUIRED COMPONENTS system thread)
find_package(nlohmann_json 3.9.1 REQUIRED)
find_package(PostgreSQL REQUIRED)
find_package(CURL REQUIRED)
find_package(OpenSSL REQUIRED)
find_package(Threads REQUIRED)

# Опция для сборки тестов
option(BUILD_TESTS "Build the tests" ON)

# Устанавливаем пути для поиска заголовочных файлов
include_directories(${PROJECT_SOURCE_DIR}/include)
include_directories(${PostgreSQL_INCLUDE_DIRS})
include_directories(${Boost_INCLUDE_DIRS})

# Собираем список исходных файлов
file(GLOB_RECURSE SOURCES "src/*.cpp")
file(GLOB_RECURSE HEADERS "include/*.hpp" "include/*.h")

# Собираем исполняемый файл
add_executable(inventory_service ${SOURCES} ${HEADERS})

# Линкуем библиотеки
target_link_libraries(inventory_service 
    PRIVATE
    ${Boost_LIBRARIES}
    nlohmann_json::nlohmann_json
    ${PostgreSQL_LIBRARIES}
    ${CURL_LIBRARIES}
    OpenSSL::SSL
    OpenSSL::Crypto
    Threads::Threads
)

# Копируем конфигурационные файлы в директорию сборки
file(COPY ${CMAKE_CURRENT_SOURCE_DIR}/config DESTINATION ${CMAKE_CURRENT_BINARY_DIR})

# Сборка тестов, если опция включена
if(BUILD_TESTS)
    add_subdirectory(test)
endif()

# Указываем директории для установки
install(TARGETS inventory_service DESTINATION bin)
install(DIRECTORY config/ DESTINATION config) 