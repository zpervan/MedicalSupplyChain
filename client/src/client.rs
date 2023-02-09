mod proto;

use druid::widget::{Flex, Label};
use druid::{AppLauncher, Data, Env, Lens, PlatformError, Widget, WindowDesc};
use futures::executor;
use grpc::ClientStubExt;
use proto::{medical_supplies::*, medical_supplies_grpc::*};

#[derive(Clone, Data, Lens)]
struct AppState {
    response: String,
}

fn main() -> Result<(), PlatformError> {
    // Example code - connecting and sending a request to the server
    let client = MedicalSuppliesClient::new_plain("127.0.0.1", 3500, Default::default()).unwrap();
    let mut req = Request::new();
    req.set_Data("Zvonimir".to_owned());

    let response = client
        .test_request(grpc::RequestOptions::new(), req)
        .join_metadata_result();

    let response_string = executor::block_on(response).unwrap().1.Data;

    // Example code - show the server response within the GUI
    let initial_state = AppState {
        response: response_string,
    };

    let main_window = WindowDesc::new(ui_builder());

    AppLauncher::with_window(main_window)
        .log_to_console()
        .launch(initial_state)
}

fn ui_builder() -> impl Widget<AppState> {
    let label = Label::new(|data: &AppState, _env: &Env| {
        if data.response.is_empty() {
            "No response yet".to_string()
        } else {
            data.response.to_string()
        }
    });

    Flex::column().with_child(label)
}
