class Favourite < ApplicationRecord
  validates :number, uniqueness: true

end
