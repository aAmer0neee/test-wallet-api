-- complex_transactions.lua
local wallet_id = "996735fb-012c-4743-8166-5bcb6039d3b3"  -- ID кошелька для тестов
local min_amount = 10    -- Минимальная сумма операции
local max_amount = 500   -- Максимальная сумма операции

-- Храним финальный баланс для проверки после теста
local balance = 1000  -- Начальное значение (лучше получать через API, но для теста фиксируем)

-- Функция выбора случайной суммы
math.randomseed(os.time())

function random_amount()
    return math.random(min_amount, max_amount)
end

-- Функция выбора случайной операции (50/50 DEPOSIT или WITHDRAW)
function random_operation()
    if math.random(0, 1) == 0 then
        return "DEPOSIT"
    else
        return "WITHDRAW"
    end
end

-- Генерация HTTP-запроса
request = function()
    local operation = random_operation()
    local amount = random_amount()

    -- Проверяем, не уходит ли баланс в минус
    if operation == "WITHDRAW" and (balance - amount < 0) then
        operation = "DEPOSIT"  -- Меняем на DEPOSIT, чтобы избежать отказов
    end

    -- Обновляем локальный баланс
    if operation == "DEPOSIT" then
        balance = balance + amount
    else
        balance = balance - amount
    end

    -- Формируем тело запроса
    local body = string.format('{"walletId": "%s", "operationType": "%s", "amount": %d}', wallet_id, operation, amount)
    
    -- Возвращаем HTTP-запрос
    return wrk.format("POST", "/api/v1/wallet", {["Content-Type"] = "application/json"}, body)
end
