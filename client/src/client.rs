mod proto;
mod gui;

use druid::{AppLauncher, Data, Lens, PlatformError, WindowDesc};
use futures::executor;
use grpc::ClientStubExt;
use proto::{medical_supplies::*, medical_supplies_grpc::*};
use spdlog::prelude::*;

#[derive(Clone, Data, Lens)]
pub struct ClientState {
    response: String,
    current_view: u32, // TODO: Add enums?
}

fn main() -> Result<(), PlatformError> {
    info!("starting client...");

    // Example code - connecting and sending a request to the server
    debug!("sending dummy request to server - will be removed later!");
    let client = MedicalSuppliesClient::new_plain("127.0.0.1", 3500, Default::default()).unwrap();
    let mut req = Request::new();
    req.set_Data("Zvonimir".to_owned());

    debug!("receiving response from server");
    let response = client
        .test_request(grpc::RequestOptions::new(), req)
        .join_metadata_result();

    let response_string = executor::block_on(response).unwrap().1.Data;

    // Example code - show the server response within the GUI
    debug!("setting initial application state");
    let initial_state = ClientState {
        response: response_string,
        current_view: 0,
    };

    let main_window = WindowDesc::new(gui::main_window::build()).menu(gui::menu_bar::make);

    debug!("starting GUI");
    AppLauncher::with_window(main_window)
        .log_to_console()
        .launch(initial_state)
}
