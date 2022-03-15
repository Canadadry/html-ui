module Main exposing (..)

import Element exposing (Element, alignRight, centerY, column, el, fill, padding, rgb255, row, spacing, text, width)
import Element.Background as Background
import Element.Border as Border
import Element.Font as Font


main =
    Element.layout []
        (column [ spacing 10 ]
            [ txt
            , el [] txt
            , el [ ft ] txt
            , row [ spacing 10, ft ] [ txt, txt ]
            ]
        )


txt : Element msg
txt =
    text "Woohoo, I'm stylish text!"


ft =
    Font.family
        [ Font.external
            { name = "Sofia"
            , url = "https://fonts.googleapis.com/css?family=Sofia"
            }
        , Font.sansSerif
        ]
