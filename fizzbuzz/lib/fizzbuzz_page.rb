require 'fizzbuzzer'

class FizzBuzzPage
  include Enumerable

  def initialize(page:, per_page:, fizzbuzzer: FizzBuzzer)
    @fizzbuzzer = fizzbuzzer
    @page = Integer(page)
    @per_page = Integer(per_page)

    @items = (start_range..end_range).map {|number| FizzBuzzer.(number)}
  end

  def each(&block)
    @items.each(&block)
  end

private

  def start_range
    (@page * @per_page) - @per_page + 1
  end

  def end_range
    range_end = @page * @per_page
  end
end
