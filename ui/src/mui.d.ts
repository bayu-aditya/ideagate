import '@mui/material/styles'
import '@mui/material/styles/createTypography'

declare module '@mui/material/styles' {
  interface PaletteOptions {
    orange?: PaletteColorOptions
    dark?: PaletteColorOptions
  }

  interface TypeText {
    dark: string
    hint: string
  }
}

declare module '@mui/material/styles/createTypography' {
  interface TypographyOptions {
    customInput: object
    mainContent: object
    menuCaption: object
    subMenuCaption: object
    commonAvatar: object
    smallAvatar: object
    mediumAvatar: object
    largeAvatar: object
  }
}
