#pragma once

#include <string>
#include <chrono>
#include <optional>
#include <vector>
#include <nlohmann/json.hpp>

namespace execution {
namespace model {

/**
 * @brief Класс, представляющий выполнение задачи ТО
 */
class TaskExecution {
public:
    TaskExecution() = default;
    ~TaskExecution() = default;

    // Getters и setters
    const std::string& getId() const { return id; }
    void setId(const std::string& id) { this->id = id; }

    const std::string& getTaskId() const { return taskId; }
    void setTaskId(const std::string& taskId) { this->taskId = taskId; }

    const std::string& getExecutorId() const { return executorId; }
    void setExecutorId(const std::string& executorId) { this->executorId = executorId; }

    const std::chrono::system_clock::time_point& getActualStartTime() const { return actualStartTime; }
    void setActualStartTime(const std::chrono::system_clock::time_point& time) { this->actualStartTime = time; }

    const std::optional<std::chrono::system_clock::time_point>& getActualEndTime() const { return actualEndTime; }
    void setActualEndTime(const std::chrono::system_clock::time_point& time) { this->actualEndTime = time; }

    const std::string& getStatus() const { return status; }
    void setStatus(const std::string& status) { this->status = status; }

    const std::string& getResult() const { return result; }
    void setResult(const std::string& result) { this->result = result; }

    const std::string& getNotes() const { return notes; }
    void setNotes(const std::string& notes) { this->notes = notes; }

    const std::string& getProblemsFound() const { return problemsFound; }
    void setProblemsFound(const std::string& problems) { this->problemsFound = problems; }

    const std::chrono::system_clock::time_point& getCreatedAt() const { return createdAt; }
    void setCreatedAt(const std::chrono::system_clock::time_point& time) { this->createdAt = time; }

    const std::chrono::system_clock::time_point& getUpdatedAt() const { return updatedAt; }
    void setUpdatedAt(const std::chrono::system_clock::time_point& time) { this->updatedAt = time; }

    // Сериализация в JSON
    nlohmann::json toJson() const;
    static TaskExecution fromJson(const nlohmann::json& json);

private:
    std::string id;                   // Уникальный идентификатор
    std::string taskId;               // Идентификатор задачи ТО
    std::string executorId;           // Идентификатор исполнителя
    std::chrono::system_clock::time_point actualStartTime; // Фактическое время начала
    std::optional<std::chrono::system_clock::time_point> actualEndTime; // Фактическое время окончания
    std::string status;               // Статус выполнения
    std::string result;               // Результат (success, failure, postponed)
    std::string notes;                // Примечания
    std::string problemsFound;        // Выявленные проблемы
    std::chrono::system_clock::time_point createdAt; // Дата создания записи
    std::chrono::system_clock::time_point updatedAt; // Дата обновления записи
};

} // namespace model
} // namespace execution 