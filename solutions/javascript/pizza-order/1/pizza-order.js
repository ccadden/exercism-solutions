/// <reference path="./global.d.ts" />
//
// @ts-check

/**
 * Determine the prize of the pizza given the pizza and optional extras
 *
 * @param {Pizza} pizza name of the pizza to be made
 * @param {Extra[]} extras list of extras
 *
 * @returns {number} the price of the pizza
 */
export function pizzaPrice(pizza, ...extras) {
  if (extras.length === 0) {
    switch (pizza) {
      case "Margherita":
        return 7
      case "Caprese":
        return 9
      case "Formaggio":
        return 10
    }
  }

  const extra = extras.pop()
  let price = 0

  switch (extra) {
    case "ExtraToppings":
      price += 2
      break
    case "ExtraSauce":
      price += 1
      break
  }

  return price + pizzaPrice(pizza, ...extras)
}

/**
 * Calculate the prize of the total order, given individual orders
 *
 * @param {PizzaOrder[]} pizzaOrders a list of pizza orders
 * @returns {number} the price of the total order
 */
export function orderPrice(pizzaOrders) {
  return pizzaOrders.reduce((acc, order) => {
    return acc + pizzaPrice(order.pizza, ...order.extras)
  }, 0)
}
