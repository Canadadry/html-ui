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
        , myCbs (\b -> Nothing)
        , myFields (\b -> Nothing)
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


myCbs : (Bool -> msg) -> Element msg
myCbs m =
    column [ spacing 30 ]
        [ Input.checkbox []
            { onChange = m
            , icon = Input.defaultCheckbox
            , checked = True
            , label =
                Input.labelRight []
                    (myElement "Do you want Guacamole? right")
            }
        , Input.checkbox []
            { onChange = m
            , icon = Input.defaultCheckbox
            , checked = True
            , label =
                Input.labelLeft []
                    (myElement "Do you want Guacamole? left ")
            }
        , Input.checkbox []
            { onChange = m
            , icon = Input.defaultCheckbox
            , checked = True
            , label =
                Input.labelAbove []
                    (myElement "Do you want Guacamole? above ")
            }
        , Input.checkbox []
            { onChange = m
            , icon = Input.defaultCheckbox
            , checked = True
            , label =
                Input.labelBelow []
                    (myElement "Do you want Guacamole? below ")
            }
        ]


myFields : (String -> msg) -> Element msg
myFields m =
    column [ spacing 30 ]
        [ Input.text []
            { onChange = m
            , text = ""
            , placeholder = Just (Input.placeholder [] (myElement "placeholder"))
            , label =
                Input.labelRight []
                    (text "Do you want Guacamole? right")
            }
        , Input.text []
            { onChange = m
            , text = "text"
            , placeholder = Just (Input.placeholder [] (myElement "placeholder"))
            , label =
                Input.labelLeft []
                    (text "Do you want Guacamole? left ")
            }
        , Input.text []
            { onChange = m
            , text = "text"
            , placeholder = Just (Input.placeholder [] (myElement "placeholder"))
            , label =
                Input.labelAbove []
                    (text "Do you want Guacamole? above ")
            }
        , Input.text []
            { onChange = m
            , text = "text"
            , placeholder = Just (Input.placeholder [] (myElement "placeholder"))
            , label =
                Input.labelBelow []
                    (text "Do you want Guacamole? below ")
            }
        ]
