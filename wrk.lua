wrk.method = "POST"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

function request()
  local userID = math.random(19, 38)
  local productID = 4
  local body = "user_id=" .. userID .. "&product_id=" .. productID
  return wrk.format(nil, nil, nil, body)
end