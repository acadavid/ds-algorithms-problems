require 'byebug'

f = File.open("input.in")
t1 = t2 = t3 = 0
f.each_line do |l|
  l.chomp!
  t1 += l.size
  t2 += eval("#{l}.size").to_i
  t3 += l.inspect.size
end

puts t1-t2 # Part 1
puts t3-t1 # Part 2
