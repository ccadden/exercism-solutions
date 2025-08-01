<?php

class PizzaPi
{
    const GRAMS_PER_PERSON = 20;
    const MIN_GRAMS = 200;

    const ML_SAUCE_PER_PIZZA = 125;

    const SLICES_PER_PIE = 8;

    public function calculateDoughRequirement($numPizzas, $numPeople)
    {
        return $numPizzas * (($numPeople * self::GRAMS_PER_PERSON) + self::MIN_GRAMS);
    }

    public function calculateSauceRequirement($numPizzas, $sauceCanVolume)
    {
        return $numPizzas * self::ML_SAUCE_PER_PIZZA / $sauceCanVolume;
    }

    public function calculateCheeseCubeCoverage($cheeseCubeEdgeLength, $cheeseLayerThickness, $pizzaDiameter)
    {
        return floor(($cheeseCubeEdgeLength**3) / ($cheeseLayerThickness * pi() * $pizzaDiameter));
    }

    public function calculateLeftOverSlices($numPizzas, $numFriends)
    {
        return ($numPizzas * self::SLICES_PER_PIE) % $numFriends;
    }
}
