def parse_steps(steps)
  floor = 0
  steps.each_char.with_index(1) do |c,i|
    if c == "("
      floor += 1
    elsif c == ")"
      floor -= 1
    end

    if floor == -1
      puts puts "Hit floor at #{i}"
    end
  end

  floor
end

f = File.open("input.in")
f.each_line do |l|
  puts parse_steps(l)
end
