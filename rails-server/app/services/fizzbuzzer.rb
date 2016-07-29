class Fizzbuzzer
  def self.call(arg)
    number = Integer(arg)

    string = ''
    string << 'Fizz' if (number % 3).zero?
    string << 'Buzz' if (number % 5).zero?

    string.empty? ? number : string
  end
end
