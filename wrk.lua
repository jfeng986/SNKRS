wrk.method = "POST"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

function request()
  local userID = math.random(132, 151)
  local productID = math.random(1, 16)
  local body = "user_id=" .. userID .. "&product_id=" .. productID
  return wrk.format(nil, nil, nil, body)
end
