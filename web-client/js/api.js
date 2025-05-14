/**
 * API модуль для взаимодействия с бекендом
 */
const API = {
    // Базовый URL API
    BASE_URL: 'http://localhost:8080',
    
    // Заголовки для запросов
    HEADERS: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
    },
    
    /**
     * Выполняет HTTP запрос
     * @param {string} endpoint - эндпоинт API
     * @param {string} method - HTTP метод (GET, POST, PUT, DELETE)
     * @param {Object} data - данные для отправки (для POST, PUT)
     * @returns {Promise} - промис с результатом запроса
     */
    async fetchAPI(endpoint, method = 'GET', data = null) {
        try {
            const url = `${this.BASE_URL}${endpoint}`;
            const options = {
                method,
                headers: this.HEADERS,
                credentials: 'include'
            };
            
            if (data && (method === 'POST' || method === 'PUT')) {
                options.body = JSON.stringify(data);
            }
            
            const response = await fetch(url, options);
            
            // Проверка на ошибки HTTP
            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(`HTTP error ${response.status}: ${errorText}`);
            }
            
            // Проверяем, есть ли контент для парсинга
            const contentType = response.headers.get('content-type');
            if (contentType && contentType.includes('application/json')) {
                return await response.json();
            } else {
                return await response.text();
            }
            
        } catch (error) {
            console.error('API request error:', error);
            throw error;
        }
    },
    
    /**
     * API для управления воздушными судами (Сервис P1)
     */
    aircraft: {
        // Получение списка ВС
        getAll: () => API.fetchAPI('/api/v1/aircraft'),
        
        // Получение информации о конкретном ВС
        getById: (id) => API.fetchAPI(`/api/v1/aircraft/${id}`),
        
        // Добавление нового ВС
        create: (data) => API.fetchAPI('/api/v1/aircraft', 'POST', data),
        
        // Обновление информации о ВС
        update: (id, data) => API.fetchAPI(`/api/v1/aircraft/${id}`, 'PUT', data),
        
        // Удаление ВС
        delete: (id) => API.fetchAPI(`/api/v1/aircraft/${id}`, 'DELETE'),
        
        // Получение типов ВС из справочника
        getTypes: () => API.fetchAPI('/api/v1/reference/aircraft-types')
    },
    
    /**
     * API для планирования ТО (Сервис P2)
     */
    maintenance: {
        // Получение всех планов ТО
        getAll: (filters = {}) => {
            const queryString = new URLSearchParams(filters).toString();
            return API.fetchAPI(`/api/v1/maintenance?${queryString}`);
        },
        
        // Получение плана ТО по ID
        getById: (id) => API.fetchAPI(`/api/v1/maintenance/${id}`),
        
        // Создание плана ТО
        create: (data) => API.fetchAPI('/api/v1/maintenance', 'POST', data),
        
        // Обновление плана ТО
        update: (id, data) => API.fetchAPI(`/api/v1/maintenance/${id}`, 'PUT', data),
        
        // Удаление плана ТО
        delete: (id) => API.fetchAPI(`/api/v1/maintenance/${id}`, 'DELETE')
    },
    
    /**
     * API для управления задачами (Сервис P3)
     */
    tasks: {
        // Получение всех задач
        getAll: (filters = {}) => {
            const queryString = new URLSearchParams(filters).toString();
            return API.fetchAPI(`/api/v1/tasks?${queryString}`);
        },
        
        // Получение задачи по ID
        getById: (id) => API.fetchAPI(`/api/v1/tasks/${id}`),
        
        // Создание задачи
        create: (data) => API.fetchAPI('/api/v1/tasks', 'POST', data),
        
        // Обновление задачи
        update: (id, data) => API.fetchAPI(`/api/v1/tasks/${id}`, 'PUT', data),
        
        // Удаление задачи
        delete: (id) => API.fetchAPI(`/api/v1/tasks/${id}`, 'DELETE'),
        
        // Изменение статуса задачи
        updateStatus: (id, status) => API.fetchAPI(`/api/v1/tasks/${id}/status`, 'PUT', { status })
    },
    
    /**
     * API для управления складом и запчастями (Сервис P4)
     */
    inventory: {
        // Получение всех запчастей
        getAll: (filters = {}) => {
            const queryString = new URLSearchParams(filters).toString();
            return API.fetchAPI(`/api/v1/inventory?${queryString}`);
        },
        
        // Получение запчасти по ID
        getById: (id) => API.fetchAPI(`/api/v1/inventory/${id}`),
        
        // Добавление новой запчасти
        create: (data) => API.fetchAPI('/api/v1/inventory', 'POST', data),
        
        // Обновление информации о запчасти
        update: (id, data) => API.fetchAPI(`/api/v1/inventory/${id}`, 'PUT', data),
        
        // Удаление запчасти
        delete: (id) => API.fetchAPI(`/api/v1/inventory/${id}`, 'DELETE'),
        
        // Создание заказа на запчасти
        createOrder: (data) => API.fetchAPI('/api/v1/inventory/orders', 'POST', data)
    },
    
    /**
     * API для аналитики и отчетов (Сервис P5)
     */
    reports: {
        // Получение отчета о состоянии парка ВС
        getFleetStatus: () => API.fetchAPI('/api/v1/reports/fleet-status'),
        
        // Получение статистики ТО
        getMaintenanceStats: (period) => API.fetchAPI(`/api/v1/reports/maintenance-stats?period=${period}`),
        
        // Получение отчета по использованию запчастей
        getPartsUsage: (filters = {}) => {
            const queryString = new URLSearchParams(filters).toString();
            return API.fetchAPI(`/api/v1/reports/parts-usage?${queryString}`);
        },
        
        // Получение отчета по эффективности персонала
        getStaffEfficiency: (filters = {}) => {
            const queryString = new URLSearchParams(filters).toString();
            return API.fetchAPI(`/api/v1/reports/staff-efficiency?${queryString}`);
        }
    }
};

// Экспортируем API для использования в других модулях
window.API = API; 