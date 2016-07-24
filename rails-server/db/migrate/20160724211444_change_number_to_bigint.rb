class ChangeNumberToBigint < ActiveRecord::Migration[5.0]
  def change
    change_column :favourites, :number, :bigint
  end
end
