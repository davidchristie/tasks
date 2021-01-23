/// <reference types="cypress" />

const ACCOUNT_MENU_SELECTOR = '[aria-controls="account-menu"]';
const LOGIN_BUTTON_TEXT = "Login";
const LOGOUT_BUTTON_TEXT = "Logout";

describe("Logging in with GitHub OAuth", () => {
  it("can log in and log out", () => {
    cy.visit("/");

    cy.contains(LOGIN_BUTTON_TEXT).click();

    cy.get(LOGIN_BUTTON_TEXT).should("not.exist");

    cy.get(ACCOUNT_MENU_SELECTOR).click();

    cy.contains(LOGOUT_BUTTON_TEXT).click();

    cy.contains(LOGIN_BUTTON_TEXT).should("exist");

    cy.get(ACCOUNT_MENU_SELECTOR).should("not.exist");
  });
});
