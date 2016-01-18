def execute(action, start_p, end_p, lights)
  case action
  when "turn on"
    res = true
  when "turn off"
    res = false
  end

  start_p[0].upto(end_p[0]) do |i|
    start_p[1].upto(end_p[1]) do |j|
      lights[i][j] = action == "toggle" ? !lights[i][j] : res
    end
  end

  total = 0
  lights.each do |i|
    total += i.count(true)
  end

  total
end

def calculate_lit_lights(lines)
  lights_lit = 0
  lights = Array.new(1000, false) { Array.new(1000, false) }
  lines.each_line do |line|
    m = /(\w* ?\w+) (\w*),(\w*) through (\w*),(\w*)/.match(line)
    action = m[1]
    pair1 = [m[2].to_i, m[3].to_i]
    pair2 = [m[4].to_i, m[5].to_i]
    lights_lit = execute(action, pair1, pair2, lights)
  end

  lights_lit
end


f = File.open("input.in")
total = calculate_lit_lights(f)
puts total
