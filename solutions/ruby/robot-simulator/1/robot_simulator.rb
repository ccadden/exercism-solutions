require 'matrix'

class Robot
  attr_reader :bearing

  DIRECTIONS = %i[north east south west].freeze
  VECTORS = [Vector[0, 1], Vector[1, 0], Vector[0, -1], Vector[-1, 0]].freeze

  BEARING_TO_VECTOR = DIRECTIONS.zip(VECTORS).to_h
  
  def initialize
    @coordinates = [0, 0]
    @direction_index = 0
  end

  def orient(direction)
    raise ArgumentError unless DIRECTIONS.include?(direction)
    
    @direction_index = DIRECTIONS.find_index(direction)
    
    @bearing = direction
  end

  def coordinates
    @coordinates.to_a
  end

  def at(x, y)
    @coordinates = Vector[x, y]
  end

  def advance
    @coordinates += BEARING_TO_VECTOR[bearing]
  end

  def turn_right
    shift_direction_index(1)
    
    @bearing = DIRECTIONS[@direction_index]
  end

  def turn_left
    shift_direction_index(-1)
    
    @bearing = DIRECTIONS[@direction_index]
  end

  private

  def shift_direction_index(shift)
    @direction_index = (@direction_index + shift) % DIRECTIONS.length
  end
end

class Simulator
  def initialize
  end

  def instructions(chars)
    res = []
    chars.each_char do |c|
      case c
        when 'L' then res << :turn_left
        when 'R' then res << :turn_right
        when 'A' then res << :advance
      end
    end

    res
  end

  def place(robot, x: 0, y: 0, direction: :north)
    robot.at(x, y)
    robot.orient(direction)
  end

  def evaluate(robot, inst)
    commands = instructions(inst)
    commands.each {|command| robot.send(command) }
  end
end