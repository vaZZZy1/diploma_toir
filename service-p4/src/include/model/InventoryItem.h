#pragma once

#include <string>
#include <chrono>
#include <optional>
#include <nlohmann/json.hpp>
#include <ctime>

namespace inventory {
namespace model {

/**
 * @brief Класс, представляющий позицию на складе
 */
class InventoryItem {
public:
    InventoryItem() = default;
    ~InventoryItem() = default;

    // Getters и setters
    const std::string& getId() const { return id; }
    void setId(const std::string& id) { this->id = id; }

    const std::string& getPartId() const { return partId; }
    void setPartId(const std::string& partId) { this->partId = partId; }

    const std::string& getWarehouseId() const { return warehouseId; }
    void setWarehouseId(const std::string& warehouseId) { this->warehouseId = warehouseId; }

    double getQuantity() const { return quantity; }
    void setQuantity(double quantity) { this->quantity = quantity; }

    const std::string& getLotNumber() const { return lotNumber; }
    void setLotNumber(const std::string& lotNumber) { this->lotNumber = lotNumber; }

    const std::string& getSerialNumber() const { return serialNumber; }
    void setSerialNumber(const std::string& serialNumber) { this->serialNumber = serialNumber; }

    const std::string& getStatus() const { return status; }
    void setStatus(const std::string& status) { this->status = status; }

    const std::string& getLocation() const { return location; }
    void setLocation(const std::string& location) { this->location = location; }

    const std::optional<std::tm>& getExpirationDate() const { return expirationDate; }
    void setExpirationDate(const std::tm& date) { this->expirationDate = date; }

    const std::string& getCertificateNumber() const { return certificateNumber; }
    void setCertificateNumber(const std::string& certificateNumber) { this->certificateNumber = certificateNumber; }

    const std::string& getPurchaseOrderId() const { return purchaseOrderId; }
    void setPurchaseOrderId(const std::string& purchaseOrderId) { this->purchaseOrderId = purchaseOrderId; }

    const std::optional<std::tm>& getArrivalDate() const { return arrivalDate; }
    void setArrivalDate(const std::tm& date) { this->arrivalDate = date; }

    const std::chrono::system_clock::time_point& getCreatedAt() const { return createdAt; }
    void setCreatedAt(const std::chrono::system_clock::time_point& time) { this->createdAt = time; }

    const std::chrono::system_clock::time_point& getUpdatedAt() const { return updatedAt; }
    void setUpdatedAt(const std::chrono::system_clock::time_point& time) { this->updatedAt = time; }

    // Сериализация в JSON
    nlohmann::json toJson() const;
    static InventoryItem fromJson(const nlohmann::json& json);

    // Проверка на просроченность
    bool isExpired() const;
    
    // Дней до истечения срока годности
    int daysUntilExpiration() const;

private:
    std::string id;                  // Уникальный идентификатор
    std::string partId;              // Запчасть
    std::string warehouseId;         // Склад
    double quantity;                 // Количество
    std::string lotNumber;           // Номер партии
    std::string serialNumber;        // Серийный номер
    std::string status;              // Статус (available, reserved, quarantine)
    std::string location;            // Местоположение на складе
    std::optional<std::tm> expirationDate; // Дата истечения срока годности
    std::string certificateNumber;   // Номер сертификата
    std::string purchaseOrderId;     // Заказ на поставку
    std::optional<std::tm> arrivalDate;    // Дата поступления
    std::chrono::system_clock::time_point createdAt; // Дата создания записи
    std::chrono::system_clock::time_point updatedAt; // Дата обновления записи
};

} // namespace model
} // namespace inventory 