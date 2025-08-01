module Port
  IDENTIFIER = :PALE

  def self.get_identifier(city)
    city.slice(0, 4).upcase.to_sym
  end

  def self.get_terminal(ship_identifier)
    code = ship_identifier.slice(0,3).upcase
    return :A if code === "OIL" || code === "GAS"

    return :B
  end
end
