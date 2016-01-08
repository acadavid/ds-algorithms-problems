def split_values(line)
  l,w,h = line.split("x")
  return l.to_i,w.to_i,h.to_i
end

def calculate_dimensions(l,w,h)
  s1 = l*w
  s2 = w*h
  s3 = h*l
  min_side = [s1,s2,s3].min
  return 2*s1 + 2*s2 + 2*s3 + min_side
end

f = File.open("input.in")
total = 0
f.each_line do |line|
  l,w,h = split_values(line)
  total += calculate_dimensions(l,w,h)
end

puts total

