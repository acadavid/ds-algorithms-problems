class String
  def is_nice?
    if (/(ab|cd|pq|xy)/ =~ self)
      return false
    end
    if !(/(.*[aeiou]){3}/ =~ self)
      return false
    end
    if !(/(.)\1{1}/ =~ self)
      return false
    end

    true
  end
end

f = File.open("input.in")
total = 0
f.each_line do |l|
  total += 1 if l.is_nice?
end

puts total

