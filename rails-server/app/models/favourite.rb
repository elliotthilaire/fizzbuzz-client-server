class Favourite < ApplicationRecord
  validates :number, uniqueness: true
  validates :number, numericality: { only_integer: true }
end
