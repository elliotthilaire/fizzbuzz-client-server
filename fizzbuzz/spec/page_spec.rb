require_relative "../lib/page"

RSpec.describe Page, "#new" do

  it 'fetches pages of items using supplied block' do
    page = Page.new(per_page: 3, starts_at:1) {|item_number| item_number * 2 }
    expect(page.to_a).to eq [2,4,6]
  end

  it 'returns array with number or per page items' do
    page = Page.new(per_page: 10) { 'item' }
    expect(page.count).to eq 10
  end

  it 'fetches the next page' do
    page = Page.new(page_number: 2, per_page: 10) { |item_number| item_number }
    expect(page.first).to eq 11
  end

  it 'is enumerable' do
  end

  context 'with default arguments' do
    it 'defaults to first page' do
      page = Page.new {'item'}
      expect(page)
    end

    it 'starts first item at 1' do
      page = Page.new { |item_number| item_number }
      expect(page.first).to eq 1
    end

    it 'defaults per_page to 100' do
      page = Page.new { 'item' }
      expect(page.count).to eq 100
    end
  end

end
