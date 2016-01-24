def look_and_say(seq)
  final_seq = []
  counter = 1
  elem = seq[0]
  seq.each_char.with_index do |o, i|
    if i+1 >= seq.size
      final_seq << counter
      final_seq << elem
      break
    end

    if seq[i] != seq[i+1]
      final_seq << counter
      final_seq << elem
      elem = seq[i+1]
      counter = 1
    else
      counter += 1
    end
  end

  final_seq.join('')
end

g_counter = 0
input = "1321131112"
while g_counter != 50 do
  input = look_and_say(input)
  g_counter += 1
end

puts input.size
