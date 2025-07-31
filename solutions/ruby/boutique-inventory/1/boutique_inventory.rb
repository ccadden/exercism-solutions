class BoutiqueInventory
  def initialize(items)
    @items = items
  end

  def item_names
    return items.map{ |item| item[:name] }.sort
  end

  def cheap
    items.filter { |item| item[:price] < 30 }
  end

  def out_of_stock
    items.filter { |item| item[:quantity_by_size].empty? }
  end

  def stock_for_item(name)
    items.find { |item| item[:name] === name }[:quantity_by_size]
  end

  def total_stock
    total = 0
    items.each { |item| item[:quantity_by_size].each { |_, v| total += v } }
    total
  end

  private
  attr_reader :items
end
