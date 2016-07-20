class Fizzbuzz
  attr_reader :number, :output, :favourite

  def initialize(number, favourites: [])
    @number = number
    @output = Fizzbuzzer.call(number)
    @favourite = favourites.include? number
  end
end
