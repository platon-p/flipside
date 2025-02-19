import styled from "styled-components";

export const Button = styled.button.attrs({
  className:
    "text-white rounded-[0.25rem] bg-primary " +
    "hover:bg-orange-600 duration-100 p-2",
})``;

export const LightButton = styled.button.attrs({
  className:
    `text-orange-600 font-medium rounded-[0.25rem] bg-orange-100
    hover:bg-orange-300 hover:text-orange-900 duration-100 p-2`,
})``;
