require 'thread'
require 'digest/md5'

input = "iwrupvqb"

def calculate(input, n, top)
  while n != top
    string = Digest::MD5.hexdigest(input+n.to_s)
    if string[0..4] == "00000"
      puts "Found: #{string} with n: #{n} \n"
      return n
    end
    n += 1
  end
  return nil
end

start_point_queue = Queue.new
(1..100000000).step(1000).to_a.each { |i| start_point_queue.push(i) }

threads_pool = (0...20).map do
  Thread.new do
      while n = start_point_queue.pop
        calculate(input, n, n+1000)
      end
  end
end

threads_pool.map(&:join)
