-- 每个元素包含一个分数和一个成员
-- 因此我们需要两个一组处理
for i = 1, #KEYS do
    local zset_key = KEYS[i]
    local score1 = ARGV[1]
    local member1 = ARGV[2]
    redis.call("ZADD", zset_key, score1, member1)
end
return "ok"