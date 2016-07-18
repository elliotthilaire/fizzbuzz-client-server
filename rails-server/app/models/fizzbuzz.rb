class Fizzbuzz
  attr_reader :number, :output

  def initialize(number)
    @number = number
    @output = Fizzbuzzer.call(number)
  end
end

