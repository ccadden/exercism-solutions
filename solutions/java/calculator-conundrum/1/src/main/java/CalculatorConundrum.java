import java.util.Set;

class CalculatorConundrum {
    private static final Set<String> VALID_OPERATIONS = Set.of("+", "*", "/");

    public String calculate(int operand1, int operand2, String operation) {
        if(operation == null) {
            throw new IllegalArgumentException("Operation cannot be null");
        }

        if (operation == "") {
            throw new IllegalArgumentException("Operation cannot be empty");
        }

        if(!isValidOperation(operation)) {
            throw new IllegalOperationException("Operation '" + operation +"' does not exist");
        }

        int answer = 0;

        switch (operation) {
            case "+":
                answer = operand1 + operand2;
                break;
            case "/":
                try {
                    answer = operand1 / operand2;
                    break;
                } catch (ArithmeticException e) {
                    throw new IllegalOperationException("Division by zero is not allowed", e);
                }
            case "*":
                answer = operand1 * operand2;
        }

        return operand1 + " " + operation + " " + operand2 + " = " + answer;
    }

    boolean isValidOperation(String operation) {
        return VALID_OPERATIONS.contains(operation);
    }
}
