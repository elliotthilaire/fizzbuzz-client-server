require 'rails_helper'

RSpec.describe Favourite, type: :model do

  it 'validates uniqness of number' do
    Favourite.create(number: 5)
    favourite = Favourite.new(number: 5)
    expect(favourite).to be_invalid
  end

  it 'validates number as integer' do
    favourite = Favourite.new(number: 'not_a_number')
    expect(favourite).to be_invalid
  end
end
