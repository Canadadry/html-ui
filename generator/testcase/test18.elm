module Main exposing (..)

import Element exposing (Element, alignRight, centerY, column, el, fill, padding, rgb255, spacing, text, width)
import Element.Background as Background
import Element.Border as Border
import Element.Font as Font
import Element.Input as Input


main =
    Element.layout []
        myRowOfStuff


myRowOfStuff =
    column [ spacing 30 ]
        [ myElement "hello"
        , myElement "hello"
        , myButton Nothing
        ]


myElement : String -> Element msg
myElement s =
    el
        [ Background.color (rgb255 240 0 245)
        , Font.color (rgb255 255 255 255)
        , Border.rounded 3
        , padding 30
        ]
        (text s)


myButton : Maybe msg -> Element msg
myButton m =
    Input.button
        [ Background.color (rgb255 238 0 0)
        , padding 30
        , Element.focused
            [ Background.color (rgb255 0 255 0) ]
        ]
        { onPress = m
        , label = myElement "My Button"
        }
