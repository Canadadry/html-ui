module Main exposing (..)

import Element exposing (Attribute, Element, alignRight, centerY, el, fill, padding, rgb255, row, spacing, text, width)
import Element.Background as Background
import Element.Border as Border
import Element.Font as Font


main =
    Element.layout []
        myRowOfStuff


myRowOfStuff : Element msg
myRowOfStuff =
    row [ width fill, centerY, spacing 30 ]
        [ myElement []
        , myElement [ shadow ]
        , el [ alignRight ] (myElement [])
        ]


myElement : List (Attribute msg) -> Element msg
myElement l =
    el
        ([ Background.color (rgb255 240 0 245)
         , Font.color (rgb255 255 255 255)
         , Border.rounded 30
         , padding 30
         ]
            ++ l
        )
        (text "stylish!")


shadow =
    Border.shadow
        { offset = ( 0, 2 )
        , size = 1
        , blur = 5
        , color = rgb255 128 128 128
        }
