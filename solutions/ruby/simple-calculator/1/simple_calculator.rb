class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze
  
  class UnsupportedOperation < StandardError
  end

  def self.calculate(first_operand, second_operand, operation)
    raise ArgumentError.new("operands must be numbers") if !first_operand.is_a?(Numeric) || !second_operand.is_a?(Numeric) 
    
    raise UnsupportedOperation if !ALLOWED_OPERATIONS.include?(operation)
    
    case operation
      when "+"
        res = first_operand + second_operand
        "#{first_operand} + #{second_operand} = #{res}"
      when "*"
        res = first_operand * second_operand
        "#{first_operand} * #{second_operand} = #{res}"
      when "/"
        res = first_operand / second_operand
        "#{first_operand} / #{second_operand} = #{res}"
    end
    rescue ZeroDivisionError
      "Division by zero is not allowed."
  end
end