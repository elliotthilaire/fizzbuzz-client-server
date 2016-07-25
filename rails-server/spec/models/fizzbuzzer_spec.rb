RSpec.describe Fizzbuzzer, "#call" do
  context "with an number divisible by 3" do
    it "returns 'Fizz'" do
      fizzbuzzed = Fizzbuzzer.(3)
      expect(fizzbuzzed).to eq 'Fizz'
    end
  end

  context "with a number divisible by 5" do
    it "returns 'Buzz'" do
      fizzbuzzed = Fizzbuzzer.(5)
      expect(fizzbuzzed).to eq 'Buzz'
    end
  end

  context "with a number divisible by 3 and 5" do
    it "returns 'FizzBuzz'" do
      fizzbuzzed = Fizzbuzzer.(3 * 5)
      expect(fizzbuzzed).to eq 'FizzBuzz'
    end
  end

  context "with a number not divisble by 3 or 5" do
    it "returns the number" do
      fizzbuzzed = Fizzbuzzer.(1)
      expect(fizzbuzzed).to eq 1
    end
  end

  context "with an argument that can be coerced to an Integer" do
    it "works" do
      fizzbuzzed = Fizzbuzzer.('7')
      expect(fizzbuzzed).to eq 7
    end
  end

  context "with an argument that cannot be coerced to a Integer" do
    it "raises an ArgumentError" do
      expect {
        Fizzbuzzer.('not_a_number')
      }.to raise_error(ArgumentError)
    end
  end
end
