
function split(str, pat)
   local t = {}
   local fpat = "(.-)" .. pat
   local last_end = 1
   local s, e, cap = str:find(fpat, 1)
   while s do
      if s ~= 1 or cap ~= "" then
     table.insert(t,cap)
      end
      last_end = e+1
      s, e, cap = str:find(fpat, last_end)
   end
   if last_end <= #str then
      cap = str:sub(last_end)
      table.insert(t, cap)
   end
   return t
end

function get_request(line)
    
    local t = {}
    local request_t = {}
    request_t.headers = {}

    local t = split(line,"\t")
    if t == nil or table.getn(t) < 20 then
        return
    end
    
    request_t.host = t[2]
    request_t.port = 8095

    local temp = split(t[7], " ")
    wrk.method = temp[1]
    wrk.path = temp[2]

    wrk.headers["upstream-status"] = t[9]
    wrk.headers["upstream-response-time"] = t[12]
    wrk.headers["bytes-sent"] = t[14]
    
    local request_length = t[13]
    wrk.body = string.rep("a", request_length)
  
    return wrk.format(nil, temp[2])
end


f = io.open("logs/access.log", "r")

request = function()
    
    local line = f:read()
    while line == nil then
        f:seek('set')
        line = f:read()
        return
    end

    return get_request(line)
end
