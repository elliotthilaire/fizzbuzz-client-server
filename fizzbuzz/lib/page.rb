
class Page
  include Enumerable

  def initialize(page: 1, per_page: 100, starts_at: 1, fetcher:)
    @page = Integer(page)
    @per_page = Integer(per_page)
    @starts_at = Integer(starts_at)

    @page_items = (start_range..end_range).map {|item_number| fetcher.(item_number) }
  end

  def each(&block)
    @page_items.each(&block)
  end

private

  def start_range
    @starts_at + (@page * @per_page) - @per_page
  end

  def end_range
    @starts_at + (@page * @per_page)
  end
end
