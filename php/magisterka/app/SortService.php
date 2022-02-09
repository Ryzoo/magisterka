<?php

namespace App;

class SortService
{
    public static function sort($elements)
    {
        $number = $elements;
        $flag = true;
        $numLength = count($number);

        for ($i = 1; $i <= $numLength - 1 && $flag; $i++) {
            $flag = false;
            for ($j = 0; $j < ($numLength - 1); $j++) {
                if ($number[$j + 1] > $number[$j]) {
                    $tmp = $number[$j];
                    $number[$j] = $number[$j + 1];
                    $number[$j + 1] = $tmp;
                    $flag = true;
                }
            }
        }

        return $number;
    }

    public static function getRandomElements(int $size)
    {
        $elements = [];

        for ($i = 0; $i < $size; $i++)
            $elements[] = rand(0, 10000);

        return $elements;
    }
}