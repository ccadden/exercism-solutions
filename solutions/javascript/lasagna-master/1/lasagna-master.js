/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

export const GRAMS_NOODLES_PER_LAYER = 50
export const LITERS_SAUCE_PER_LAYER = 0.2

export function cookingStatus(remainingTime) {
  switch (remainingTime) {
    case undefined:
      return 'You forgot to set the timer.'
    case 0:
      return 'Lasagna is done.'
    default:
      return 'Not done, please wait.'
  }
}

export function preparationTime(layers, prepTimePerLayer = 2) {
  return prepTimePerLayer * layers.length
}

export function quantities(layers) {
  let obj = {
    noodles: 0,
    sauce: 0,
  }

  layers.forEach((layer) => {
    switch (layer) {
      case "noodles":
        obj.noodles += GRAMS_NOODLES_PER_LAYER
        break
      case "sauce":
        obj.sauce += LITERS_SAUCE_PER_LAYER
    }
  })

  return obj
}

export function addSecretIngredient(arr1, arr2) {
  arr2.push(...arr1.slice(-1))
}

export function scaleRecipe(recipe, numPortions) {
  const obj = {}
  for (let key in recipe) {
    obj[key] = recipe[key] * numPortions / 2
  }

  return obj
}
