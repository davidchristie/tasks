/// <reference path="../support/index.d.ts" />

describe("Login", () => {
  const ACCOUNT_MENU = "account-menu";
  const LOGIN_BUTTON = "topbar-login";
  const LOGOUT_BUTTON = "account-menu-logout";

  it("should login when the 'Login' button is clicked", () => {
    cy.visit("/");

    cy.getDataTest(LOGIN_BUTTON).click();

    cy.getDataTest(LOGIN_BUTTON).should("not.exist");

    cy.getDataTest(ACCOUNT_MENU).should("exist");
  });

  it("should logout when the 'Logout' button is clicked", () => {
    cy.login();

    cy.getDataTest(ACCOUNT_MENU).click();

    cy.getDataTest(LOGOUT_BUTTON).click();

    cy.getDataTest(LOGIN_BUTTON).should("exist");

    cy.getDataTest(ACCOUNT_MENU).should("not.exist");
  });
});
