class AssemblyLine
  BASE_CARS_PER_HOUR = 221
  MINUTES_PER_HOUR = 60
  
  def initialize(speed)
    @speed = speed
  end

  def success_rate
    case @speed
      when 1..4
        1.0
      when 5..8
        0.9
      when 9
        0.8
      else
        0.77
    end
  end

  def production_rate_per_hour
    @speed * success_rate * BASE_CARS_PER_HOUR
  end

  def working_items_per_minute
    (production_rate_per_hour / MINUTES_PER_HOUR).floor
  end
end
