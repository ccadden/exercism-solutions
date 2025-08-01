class Matrix
  attr_reader :matrix
  def initialize(matrix)
    @matrix = parse_matrix(matrix)
  end

  def parse_matrix(matrix)
    rows = matrix.split("\n").map{ |row| row.split(" ").map(&:to_i) }
  end

  def column(n)
    throw ArgumentError.new() if n < 1
    matrix.map { |row| row[n-1] }
  end

  def row(n)
    throw ArgumentError.new() if n < 1
    return matrix[n-1]
  end
end