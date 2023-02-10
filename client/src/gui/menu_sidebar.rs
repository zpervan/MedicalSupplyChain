use druid::{Env, Widget, WidgetExt};
use druid::widget::{Button, Flex, Label, ViewSwitcher};
use crate::ClientState;

const BUTTON_SIZE: f64 = 140_f64;

pub fn make() -> impl Widget<ClientState> {
    let mut sidebar = Flex::column();

    sidebar.add_child(Label::new(|data: &u32, _env: &Env| format!("Current view: {data}")).lens(ClientState::current_view));

    // TODO: Create factory function which will add a new button
    sidebar.add_spacer(80.);

    // Add "Inventory button"
    sidebar.add_child(Button::new("Inventory").on_click(move |_event, data: &mut u32, _env| {
        println!("button \"Inventory\" is clicked");
        *data = 0;
    })
        .fix_size(BUTTON_SIZE, BUTTON_SIZE)
        .lens(ClientState::current_view));

    sidebar.add_spacer(80.);

    // Add "Users" button
    sidebar.add_child(Button::new("Users").on_click(move |_event, data: &mut u32, _env| {
        println!("button \"Users\" is clicked");
        *data = 1;
    })
        .fix_size(BUTTON_SIZE, BUTTON_SIZE)
        .lens(ClientState::current_view));

    let view_switcher = ViewSwitcher::new(
        |data: &ClientState, _env| data.current_view,
        |selector, _data, _env| match selector {
            0 => Box::new(Label::new("Inventory view").center()),
            1 => Box::new(Label::new("Users view").center()),
            _ => Box::new(Label::new("Unknown").center())
        });

    Flex::row()
        .with_child(sidebar)
        .with_flex_child(view_switcher, 1.0)
}