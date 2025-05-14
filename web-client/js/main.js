document.addEventListener('DOMContentLoaded', function() {
    initBootstrapComponents();
    initNavigation();
    initModalHandlers();
    initFormHandlers();
    initFiltersAndSearch();
    initReportHandlers();
});

function initBootstrapComponents() {
    const tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
    tooltipTriggerList.map(tooltipTriggerEl => new bootstrap.Tooltip(tooltipTriggerEl));
    
    const popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'));
    popoverTriggerList.map(popoverTriggerEl => new bootstrap.Popover(popoverTriggerEl));
}

function initNavigation() {
    const navLinks = document.querySelectorAll('.nav-link[data-section]');
    
    navLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            
            navLinks.forEach(l => l.classList.remove('active'));
            
            this.classList.add('active');
            
            const sectionId = this.getAttribute('data-section');
            
            document.querySelectorAll('.content-section').forEach(section => {
                section.classList.add('d-none');
            });
            
            document.getElementById('welcome-section').classList.add('d-none');
            
            const targetSection = document.getElementById(`${sectionId}-section`);
            if (targetSection) {
                targetSection.classList.remove('d-none');
                
                loadSectionData(sectionId);
            }
        });
    });
}

function loadSectionData(sectionId) {
    switch(sectionId) {
        case 'aircraft':
            loadAircraftData();
            break;
        case 'maintenance':
            loadMaintenanceData();
            break;
        case 'tasks':
            loadTasksData();
            break;
        case 'inventory':
            loadInventoryData();
            break;
        case 'reports':
            break;
    }
}

async function loadAircraftData() {
    try {
        const aircraftTableBody = document.getElementById('aircraft-table-body');
        aircraftTableBody.innerHTML = '<tr><td colspan="5" class="text-center">Загрузка данных...</td></tr>';
        
        const data = await API.aircraft.getAll();
        
        if (data && data.length > 0) {
            let html = '';
            
            data.forEach(aircraft => {
                html += `
                <tr data-id="${aircraft.id}">
                    <td>${aircraft.registration_number}</td>
                    <td>${aircraft.type}</td>
                    <td>${aircraft.serial_number}</td>
                    <td>${aircraft.year}</td>
                    <td>
                        <button class="btn btn-sm btn-outline-primary action-btn edit-aircraft-btn" data-id="${aircraft.id}">
                            <i class="bi bi-pencil"></i> Изменить
                        </button>
                        <button class="btn btn-sm btn-outline-danger action-btn delete-aircraft-btn" data-id="${aircraft.id}">
                            <i class="bi bi-trash"></i> Удалить
                        </button>
                    </td>
                </tr>
                `;
            });
            
            aircraftTableBody.innerHTML = html;
            
            document.querySelectorAll('.edit-aircraft-btn').forEach(btn => {
                btn.addEventListener('click', function() {
                    const aircraftId = this.getAttribute('data-id');
                    openAircraftEditModal(aircraftId);
                });
            });
            
            document.querySelectorAll('.delete-aircraft-btn').forEach(btn => {
                btn.addEventListener('click', function() {
                    const aircraftId = this.getAttribute('data-id');
                    deleteAircraft(aircraftId);
                });
            });
            
        } else {
            aircraftTableBody.innerHTML = '<tr><td colspan="5" class="text-center">Нет данных о воздушных судах</td></tr>';
        }
        
        loadAircraftTypes();
        
    } catch (error) {
        console.error('Error loading aircraft data:', error);
        document.getElementById('aircraft-table-body').innerHTML = 
            '<tr><td colspan="5" class="text-center text-danger">Ошибка загрузки данных</td></tr>';
    }
}

async function loadAircraftTypes() {
    try {
        const aircraftTypeSelect = document.getElementById('aircraft-type');
        
        while (aircraftTypeSelect.options.length > 0) {
            aircraftTypeSelect.remove(0);
        }
        
        const defaultOption = document.createElement('option');
        defaultOption.value = '';
        defaultOption.textContent = 'Выберите тип ВС';
        defaultOption.selected = true;
        defaultOption.disabled = true;
        aircraftTypeSelect.appendChild(defaultOption);
        
        const types = await API.aircraft.getTypes();
        
        if (types && types.length > 0) {
            types.forEach(type => {
                const option = document.createElement('option');
                option.value = type.id;
                option.textContent = type.name;
                aircraftTypeSelect.appendChild(option);
            });
        }
        
    } catch (error) {
        console.error('Error loading aircraft types:', error);
    }
}

async function openAircraftEditModal(aircraftId) {
    try {
        const modal = new bootstrap.Modal(document.getElementById('aircraft-modal'));
        document.querySelector('#aircraft-modal .modal-title').textContent = 'Редактирование воздушного судна';
        
        const aircraft = await API.aircraft.getById(aircraftId);
        
        document.getElementById('aircraft-reg').value = aircraft.registration_number;
        document.getElementById('aircraft-serial').value = aircraft.serial_number;
        document.getElementById('aircraft-year').value = aircraft.year;
        
        const typeSelect = document.getElementById('aircraft-type');
        for (let i = 0; i < typeSelect.options.length; i++) {
            if (typeSelect.options[i].value === aircraft.type_id) {
                typeSelect.selectedIndex = i;
                break;
            }
        }
        
        document.getElementById('aircraft-form').setAttribute('data-id', aircraftId);
        
        modal.show();
        
    } catch (error) {
        console.error('Error opening aircraft edit modal:', error);
        alert('Не удалось загрузить данные для редактирования');
    }
}

async function deleteAircraft(aircraftId) {
    if (confirm('Вы уверены, что хотите удалить это воздушное судно?')) {
        try {
            await API.aircraft.delete(aircraftId);
            loadAircraftData();
        } catch (error) {
            console.error('Error deleting aircraft:', error);
            alert('Не удалось удалить воздушное судно');
        }
    }
}

async function loadMaintenanceData() {
    try {
        const maintenanceCalendar = document.getElementById('maintenance-calendar');
        maintenanceCalendar.innerHTML = '<div class="text-center p-5">Загрузка данных...</div>';
        
        const filter = document.getElementById('maintenance-filter').value;
        const date = document.getElementById('maintenance-date').value;
        
        const params = {};
        if (filter && filter !== 'all') {
            params.status = filter;
        }
        if (date) {
            params.date = date;
        }
        
        const data = await API.maintenance.getAll(params);
        
        if (data && data.length > 0) {
            let html = '<div class="list-group">';
            
            data.forEach(maintenance => {
                const statusClass = getStatusClass(maintenance.status);
                
                html += `
                <a href="#" class="list-group-item list-group-item-action" data-id="${maintenance.id}">
                    <div class="d-flex w-100 justify-content-between">
                        <h5 class="mb-1">${maintenance.aircraft.registration_number} - ${maintenance.title}</h5>
                        <span class="status-badge ${statusClass}">${maintenance.status}</span>
                    </div>
                    <p class="mb-1">${maintenance.description}</p>
                    <small>Запланировано на: ${new Date(maintenance.scheduled_date).toLocaleDateString()}</small>
                </a>
                `;
            });
            
            html += '</div>';
            maintenanceCalendar.innerHTML = html;
            
            document.querySelectorAll('#maintenance-calendar .list-group-item').forEach(item => {
                item.addEventListener('click', function(e) {
                    e.preventDefault();
                    const maintenanceId = this.getAttribute('data-id');
                    openMaintenanceDetailsModal(maintenanceId);
                });
            });
            
        } else {
            maintenanceCalendar.innerHTML = '<div class="text-center p-5">Нет данных о планах ТО</div>';
        }
        
    } catch (error) {
        console.error('Error loading maintenance data:', error);
        document.getElementById('maintenance-calendar').innerHTML = 
            '<div class="text-center p-5 text-danger">Ошибка загрузки данных</div>';
    }
}

async function loadTasksData() {
    try {
        const tasksTableBody = document.getElementById('tasks-table-body');
        tasksTableBody.innerHTML = '<tr><td colspan="7" class="text-center">Загрузка данных...</td></tr>';
        
        const filter = document.getElementById('tasks-filter').value;
        
        const params = {};
        if (filter && filter !== 'all') {
            params.status = filter;
        }
        
        const data = await API.tasks.getAll(params);
        
        if (data && data.length > 0) {
            let html = '';
            
            data.forEach(task => {
                const statusClass = getStatusClass(task.status);
                
                html += `
                <tr data-id="${task.id}">
                    <td>${task.id}</td>
                    <td>${task.description}</td>
                    <td>${task.aircraft.registration_number}</td>
                    <td><span class="status-badge ${statusClass}">${task.status}</span></td>
                    <td>${task.assigned_to || '-'}</td>
                    <td>${new Date(task.due_date).toLocaleDateString()}</td>
                    <td>
                        <button class="btn btn-sm btn-outline-primary action-btn edit-task-btn" data-id="${task.id}">
                            <i class="bi bi-pencil"></i>
                        </button>
                        <button class="btn btn-sm btn-outline-success action-btn status-task-btn" data-id="${task.id}">
                            <i class="bi bi-check-circle"></i>
                        </button>
                        <button class="btn btn-sm btn-outline-danger action-btn delete-task-btn" data-id="${task.id}">
                            <i class="bi bi-trash"></i>
                        </button>
                    </td>
                </tr>
                `;
            });
            
            tasksTableBody.innerHTML = html;
            
            document.querySelectorAll('.edit-task-btn').forEach(btn => {
                btn.addEventListener('click', function() {
                    const taskId = this.getAttribute('data-id');
                });
            });
            
            document.querySelectorAll('.status-task-btn').forEach(btn => {
                btn.addEventListener('click', function() {
                    const taskId = this.getAttribute('data-id');
                });
            });
            
            document.querySelectorAll('.delete-task-btn').forEach(btn => {
                btn.addEventListener('click', function() {
                    const taskId = this.getAttribute('data-id');
                });
            });
            
        } else {
            tasksTableBody.innerHTML = '<tr><td colspan="7" class="text-center">Нет данных о задачах</td></tr>';
        }
        
    } catch (error) {
        console.error('Error loading tasks data:', error);
        document.getElementById('tasks-table-body').innerHTML = 
            '<tr><td colspan="7" class="text-center text-danger">Ошибка загрузки данных</td></tr>';
    }
}

async function loadInventoryData() {
    try {
        const inventoryTableBody = document.getElementById('inventory-table-body');
        inventoryTableBody.innerHTML = '<tr><td colspan="7" class="text-center">Загрузка данных...</td></tr>';
        
        const searchQuery = document.getElementById('inventory-search').value;
        
        const params = {};
        if (searchQuery) {
            params.search = searchQuery;
        }
        
        const data = await API.inventory.getAll(params);
        
        if (data && data.length > 0) {
            let html = '';
            
            data.forEach(part => {
                const statusClass = part.quantity > part.min_quantity ? 'status-completed' : 'status-critical';
                const statusText = part.quantity > part.min_quantity ? 'В наличии' : 'Требуется заказ';
                
                html += `
                <tr data-id="${part.id}">
                    <td>${part.id}</td>
                    <td>${part.name}</td>
                    <td>${part.part_number}</td>
                    <td>${part.quantity}</td>
                    <td>${part.location || '-'}</td>
                    <td><span class="status-badge ${statusClass}">${statusText}</span></td>
                    <td>
                        <button class="btn btn-sm btn-outline-primary action-btn edit-part-btn" data-id="${part.id}">
                            <i class="bi bi-pencil"></i>
                        </button>
                        <button class="btn btn-sm btn-outline-danger action-btn delete-part-btn" data-id="${part.id}">
                            <i class="bi bi-trash"></i>
                        </button>
                    </td>
                </tr>
                `;
            });
            
            inventoryTableBody.innerHTML = html;
            
            document.querySelectorAll('.edit-part-btn').forEach(btn => {
                btn.addEventListener('click', function() {
                    const partId = this.getAttribute('data-id');
                });
            });
            
            document.querySelectorAll('.delete-part-btn').forEach(btn => {
                btn.addEventListener('click', function() {
                    const partId = this.getAttribute('data-id');
                });
            });
            
        } else {
            inventoryTableBody.innerHTML = '<tr><td colspan="7" class="text-center">Нет данных о запчастях</td></tr>';
        }
        
    } catch (error) {
        console.error('Error loading inventory data:', error);
        document.getElementById('inventory-table-body').innerHTML = 
            '<tr><td colspan="7" class="text-center text-danger">Ошибка загрузки данных</td></tr>';
    }
}

function initFormHandlers() {
    document.getElementById('save-aircraft-btn')?.addEventListener('click', async function() {
        const form = document.getElementById('aircraft-form');
        
        if (!form.checkValidity()) {
            form.reportValidity();
            return;
        }
        
        const data = {
            registration_number: document.getElementById('aircraft-reg').value,
            type_id: document.getElementById('aircraft-type').value,
            serial_number: document.getElementById('aircraft-serial').value,
            year: parseInt(document.getElementById('aircraft-year').value)
        };
        
        try {
            const aircraftId = form.getAttribute('data-id');
            
            if (aircraftId) {
                await API.aircraft.update(aircraftId, data);
            } else {
                await API.aircraft.create(data);
            }
            
            const modal = bootstrap.Modal.getInstance(document.getElementById('aircraft-modal'));
            modal.hide();
            
            loadAircraftData();
            
        } catch (error) {
            console.error('Error saving aircraft:', error);
            alert('Не удалось сохранить данные');
        }
    });
    
    document.getElementById('add-aircraft-btn')?.addEventListener('click', function() {
        document.getElementById('aircraft-form').reset();
        document.getElementById('aircraft-form').removeAttribute('data-id');
        
        document.querySelector('#aircraft-modal .modal-title').textContent = 'Добавление воздушного судна';
        
        const modal = new bootstrap.Modal(document.getElementById('aircraft-modal'));
        modal.show();
    });
}

function initFiltersAndSearch() {
    document.getElementById('tasks-filter')?.addEventListener('change', function() {
        loadTasksData();
    });
    
    document.getElementById('maintenance-filter')?.addEventListener('change', function() {
        loadMaintenanceData();
    });
    
    document.getElementById('maintenance-date')?.addEventListener('change', function() {
        loadMaintenanceData();
    });
    
    document.getElementById('inventory-search')?.addEventListener('input', debounce(function() {
        loadInventoryData();
    }, 500));
    
    document.getElementById('aircraft-search')?.addEventListener('input', debounce(function() {
        loadAircraftData();
    }, 500));
}

function initReportHandlers() {
    document.querySelectorAll('.report-btn')?.forEach(btn => {
        btn.addEventListener('click', async function() {
            const reportType = this.getAttribute('data-report');
            const reportContent = document.getElementById('report-content');
            
            reportContent.innerHTML = '<div class="text-center p-5">Загрузка отчета...</div>';
            
            try {
                let data;
                
                switch(reportType) {
                    case 'fleet-status':
                        data = await API.reports.getFleetStatus();
                        displayFleetStatusReport(data);
                        break;
                    case 'maintenance-stats':
                        data = await API.reports.getMaintenanceStats('month');
                        displayMaintenanceStatsReport(data);
                        break;
                    case 'parts-usage':
                        data = await API.reports.getPartsUsage();
                        displayPartsUsageReport(data);
                        break;
                    case 'staff-efficiency':
                        data = await API.reports.getStaffEfficiency();
                        displayStaffEfficiencyReport(data);
                        break;
                }
                
            } catch (error) {
                console.error('Error loading report:', error);
                reportContent.innerHTML = '<div class="alert alert-danger">Ошибка при загрузке отчета</div>';
            }
        });
    });
}

function displayFleetStatusReport(data) {
    const reportContent = document.getElementById('report-content');
    
    let html = '<h3>Состояние парка воздушных судов</h3>';
    
    html += `
    <div class="row mt-4">
        <div class="col-md-3">
            <div class="card bg-success text-white">
                <div class="card-body text-center">
                    <h5 class="card-title">Исправные ВС</h5>
                    <h2>${data.operational}</h2>
                </div>
            </div>
        </div>
        <div class="col-md-3">
            <div class="card bg-warning text-dark">
                <div class="card-body text-center">
                    <h5 class="card-title">На ТО</h5>
                    <h2>${data.maintenance}</h2>
                </div>
            </div>
        </div>
        <div class="col-md-3">
            <div class="card bg-danger text-white">
                <div class="card-body text-center">
                    <h5 class="card-title">Неисправные</h5>
                    <h2>${data.inoperative}</h2>
                </div>
            </div>
        </div>
        <div class="col-md-3">
            <div class="card bg-info text-white">
                <div class="card-body text-center">
                    <h5 class="card-title">Всего</h5>
                    <h2>${data.total}</h2>
                </div>
            </div>
        </div>
    </div>
    `;
    
    reportContent.innerHTML = html;
}

function initModalHandlers() {
    document.querySelectorAll('.modal').forEach(modal => {
        modal.addEventListener('hidden.bs.modal', function() {
            const forms = this.querySelectorAll('form');
            forms.forEach(form => {
                form.reset();
                form.removeAttribute('data-id');
            });
        });
    });
}

function getStatusClass(status) {
    switch(status.toLowerCase()) {
        case 'pending':
        case 'scheduled':
            return 'status-pending';
        case 'in progress':
        case 'in-progress':
            return 'status-inprogress';
        case 'completed':
        case 'done':
            return 'status-completed';
        case 'critical':
        case 'overdue':
            return 'status-critical';
        default:
            return '';
    }
}

function debounce(func, wait) {
    let timeout;
    return function() {
        const context = this;
        const args = arguments;
        clearTimeout(timeout);
        timeout = setTimeout(() => {
            func.apply(context, args);
        }, wait);
    };
}