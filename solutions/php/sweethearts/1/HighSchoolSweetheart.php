<?php

class HighSchoolSweetheart
{
    public function firstLetter(string $name): string
    {
        $name = trim($name);
        return$name[0];
    }

    public function initial(string $name): string
    {
        return strtoupper(self::firstLetter($name)) . ".";
    }

    public function initials(string $name): string
    {
        $names = explode(" ", $name);
        return self::initial($names[0]) . " " . self::initial($names[1]);
    }

    public function pair(string $sweetheart_a, string $sweetheart_b): string
    {
        $inits1 = self::initials($sweetheart_a);
        $inits2 = self::initials($sweetheart_b);
        return <<<HEART
     ******       ******
   **      **   **      **
 **         ** **         **
**            *            **
**                         **
**     $inits1  +  $inits2     **
 **                       **
   **                   **
     **               **
       **           **
         **       **
           **   **
             ***
              *
HEART;
    }
}
