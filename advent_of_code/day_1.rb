def parse_steps(steps)
  floor = 0
  steps.each_char do |c|
    if c == "("
      floor += 1
    elsif c == ")"
      floor -= 1
    end
  end

  floor
end

f = File.open("input.in")
f.each_line do |l|
  puts parse_steps(l)
end
