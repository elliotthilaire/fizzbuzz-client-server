class Page
  include Enumerable

  attr_reader :page, :per_page

  def initialize(page: 1, per_page: 100, starts_at: 1, &block)
    @page = Integer(page)
    @per_page = Integer(per_page)
    @starts_at = Integer(starts_at)

    @page_items = item_numbers.map {|item_number| yield(item_number) }
  end

  def each(&block)
    @page_items.each(&block)
  end

private

  def item_numbers
    (start_range..end_range)
  end

  def start_range
    @starts_at + (@page * @per_page) - @per_page
  end

  def end_range
    @starts_at + (@page * @per_page) - 1
  end
end
