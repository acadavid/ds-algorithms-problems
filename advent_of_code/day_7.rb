require 'byebug'

def find_value(obj, file)
  file.each_line do |l|
    case l
    when /^(\d+) -> #{obj}$/
      file.rewind
      return $1
    when /(.*) -> #{obj}$/
      file.rewind
      return $1
    end
  end
end

def parse(val)
  case val
  when /NOT (.*)$/
    [:~, $1, nil]
  when /(.*) AND (.*)/
    [:&, $1, $2]
  when /(.*) OR (.*)/
    [:|, $1, $2]
  when /(.*) RSHIFT (.*)/
    [:>>, $1, $2]
  when /(.*) LSHIFT (.*)/
    [:<<, $1, $2]
  when /(.*)/
    [nil, val, nil]
  end
end

def find_output(obj, file, memo={})
  if (/\A[-+]?[0-9]+\z/ =~ obj) != nil
    return obj.to_i
  end
  val = find_value(obj, file)
  op, fst, snd = parse(val)

  if op.nil?
    return find_output(fst, file, memo)
  elsif op == :~
    return eval("#{op} #{find_output(fst, file, memo)}")
  else
    fst_res = memo[fst] || find_output(fst, file, memo)
    snd_res = memo[snd] || find_output(snd, file, memo)
    res = eval("#{fst_res} #{op} #{snd_res}")
    memo[obj] = res
    return res
  end
end

file = File.open('input.in')
puts find_output("a", file)

