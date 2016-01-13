require 'set'

def step(m,pos)
  case m
  when ">" then pos[0] += 1
  when "<" then pos[0] -= 1
  when "^" then pos[1] += 1
  when "v" then pos[1] -= 1
  end
  pos
end

def calculate_visited_places(line)
  visited = Set.new [[0,0]]
  r = [0,0]
  s = [0,0]

  line.each_char.with_index do |m,i|
    if i.even?
      l = step(m,s)
      visited << l.dup
      s = l
    else
      l = step(m,r)
      visited << l.dup
      r = l
    end
  end

  return visited.size
end

f = File.open("input.in")
f.each_line do |l|
  puts calculate_visited_places(l)
end

