class LocomotiveEngineer
  def self.generate_list_of_wagons(*args)
    return *args
  end

  def self.fix_list_of_wagons(each_wagons_id, missing_wagons)
    misplaced_wagon1, misplaced_wagon2, first_wagon, *remaining_wagons = each_wagons_id
    [first_wagon, *missing_wagons, *remaining_wagons, misplaced_wagon1, misplaced_wagon2]
  end

  def self.add_missing_stops(routing_hash, **stops)
    routing_hash[:stops] = stops.values
    routing_hash
  end

  def self.extend_route_information(route, more_route_information)
    {**route, **more_route_information}
  end
end
