#TSP so let's brute

def find_paths(cities)
  totals = []
  cities.keys.permutation.each do |perm|
    d = 0
    perm.each_cons(2) do |pp|
      d += cities[pp[0]][pp[1]]
    end
    totals << d
  end

  totals
end

cities = Hash.new{ |h,k| h[k] = {} }
f = File.open("input.in")
f.each_line do |l|
  m = /(.*) to (.*) = (\d*)/.match(l)
  s = m[1]
  d = m[2]
  w = m[3].to_i
  cities[s][d] = w
  cities[d][s] = w
end

totals = find_paths(cities)
puts totals.min # Part 1
puts totals.max # Part 2
