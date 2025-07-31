require 'date'

class Meetup
  attr_reader :days_in_month
  
  def initialize(month, year)
    @days_in_month = Date.new(year, month, 1)..Date.new(year, month, -1)
  end

  def day(name, ordinal)
    days = days_in_month.filter(&:"#{name}?")
    
    case ordinal
      when :teenth then days.find { |day| day.day.between?(13, 19)}
      when :first then days[0]
      when :second then days[1]
      when :third then days[2]
      when :fourth then days[3]
      when :last then days[-1]
    end
  end
end