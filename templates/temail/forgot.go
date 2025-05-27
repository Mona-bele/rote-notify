package temail

const FORGOT_PASSWORD = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>RESET PASSWORD</title>
    <style>
        body {
            padding: 0;
            margin: 0;
            font-family: Roboto, Arial, sans-serif;
            background: #fff;
        }
        .ger-top span {
            font-size: 14px;
            color: #575757;
        }
        .ger-text {
            color: #ffffff;
            margin: 10rem auto;
            width: 80%;
            font-size: 1.50rem;
        }
        .ger-container {
            margin: 5rem auto;
            width: 500px;
            text-align: center;
        }
        .ger-top {
            line-height: 15px;
        }
        .ger-inputs input {
            margin-top: 30px;
            padding: 0.98rem;
            width: 225px;
            height: 15px;
            border: 2px solid transparent;
            border-radius: 4px;
            outline: 0;
            background: #f7f7f7;
            margin-bottom: 15px;
        }
        .ger-inputs input:focus {
            border: 2px solid #FF2D9E !important;
            padding: 0.98rem;
            width: 100%;
            height: 20px;
            border: 0;
            border-radius: 4px;
            transition: 2s all;
            color: #7300E6;
        }
        .ger-btn button {
            background: #e60086;
            border: 0;
            width: 260px;
            padding: 0.98%;
            height: 52px;
            color: #fff;
            font-size: 14px;
            border-radius: 4px;
            cursor: pointer;
        }
    </style>
</head>
<body style="padding: 0; margin: 0; font-family: Roboto, Arial, sans-serif; background-color: #fff;">
<table role="presentation" cellspacing="0" cellpadding="0" border="0" align="center" width="100%" style="max-width: 600px; margin: auto; border-collapse: collapse;">
    <tr>
        <td style="padding: 20px 0; text-align: center;">
            <svg width="46" height="65" viewBox="0 0 46 65" fill="none" xmlns="http://www.w3.org/2000/svg" style="display: block; margin: auto;">
                <path d="M10.0158 65L6.58008 56.8993L18.272 51.932V13.7462L6.78011 7.82407L10.8049 0L27.0645 8.37849V57.7568L10.0158 65Z" fill="black"/>
                <path d="M8.79255 18.4054H0V46.4566H8.79255V18.4054Z" fill="black"/>
                <path d="M28.2877 65L24.8521 56.8993L36.5439 51.932V13.7462L25.0521 7.82407L29.0769 0L45.3365 8.37849V57.7568L28.2877 65Z" fill="black"/>
            </svg>
            <h1 style="color: #1E1E1E; font-size: 28px; font-weight: bold;">
                Hello <span style="color: #e60086;">!NAME</span>, 👋🏽</h1>
            <p style="color: #575757; font-size: 14px;">Forgot your password?</p>
            <p style="color: #575757; font-size: 14px;">Below is your password reset code:</p>
        </td>
    </tr>
    <tr>
        <td style="text-align: center; padding: 100px 10px;">
            <table role="presentation" cellspacing="0" cellpadding="0" border="0" align="center" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;">
                <tr>
                    <td style="border-radius: 4px; background-color: #f7f7f7; padding: 10px;">
                        <div style="width: 100%; border: 2px solid transparent; border-radius: 4px; padding: 10px;">
                            <!--<p style="color: #575757; font-size: 14px; margin: 0;">Your verification code is:</p>-->
                            <h2 style="color: #1E1E1E; font-size: 24px; font-weight: bold; margin: 0; text-align: center">!CODE</h2>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td style="padding: 20px 0;">
                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" align="center" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
                            <tr>
                                <!--<td style="border-radius: 4px; background-color: #e60086;">
                                    <a href="#" style="color: #ffffff; text-decoration: none; font-size: 14px; display: inline-block; padding: 15px 30px; border-radius: 4px;">Confirm</a>
                                </td>-->
                            </tr>
                        </table>
                    </td>
                </tr>
            </table>
        </td>
    </tr>
</table>
</body>
</html>
`
