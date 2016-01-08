require 'set'

def calculate_visited_places(line)
  x = y = 0
  visited = Set.new [[0,0]]
  line.each_char do |m|
    case m
    when ">"
      x += 1
    when "<"
      x -= 1
    when "^"
      y += 1
    when "v"
      y -= 1
    end
    visited << [x,y]
  end

  return visited.size
end

f = File.open("input.in")
f.each_line do |l|
  puts calculate_visited_places(l)
end
